package connector

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/connectivity"
)

// HealthChecker 健康检查器
type HealthChecker struct {
	manager  *Manager
	interval time.Duration
	debug    bool
	stopCh   chan struct{}
	wg       sync.WaitGroup
	running  bool
	mu       sync.Mutex
}

// NewHealthChecker 创建新的健康检查器
func NewHealthChecker(manager *Manager, interval time.Duration, debug bool) *HealthChecker {
	return &HealthChecker{
		manager:  manager,
		interval: interval,
		debug:    debug,
		stopCh:   make(chan struct{}),
	}
}

// Start 启动健康检查
func (hc *HealthChecker) Start() {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	if hc.running {
		return
	}

	hc.running = true
	hc.wg.Add(1)

	go hc.run()

	if hc.debug {
		log.Printf("[INFO] HealthChecker.Start: 健康检查已启动，间隔: %v", hc.interval)
	}
}

// Stop 停止健康检查
func (hc *HealthChecker) Stop() {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	if !hc.running {
		return
	}

	hc.running = false
	close(hc.stopCh)
	hc.wg.Wait()

	if hc.debug {
		log.Printf("[INFO] HealthChecker.Stop: 健康检查已停止")
	}
}

// run 运行健康检查循环
func (hc *HealthChecker) run() {
	defer hc.wg.Done()

	ticker := time.NewTicker(hc.interval)
	defer ticker.Stop()

	for {
		select {
		case <-hc.stopCh:
			return
		case <-ticker.C:
			hc.checkHealth()
		}
	}
}

// checkHealth 执行健康检查
func (hc *HealthChecker) checkHealth() {
	if hc.manager == nil {
		return
	}

	state := hc.manager.GetState()

	if hc.debug {
		log.Printf("[DEBUG] HealthChecker.checkHealth: 当前连接状态: %v", state)
	}

	// 如果连接状态不健康，尝试重连
	switch state {
	case connectivity.TransientFailure, connectivity.Shutdown:
		if hc.debug {
			log.Printf("[WARN] HealthChecker.checkHealth: 检测到不健康状态 %v，尝试重连", state)
		}
		hc.attemptReconnect()
	case connectivity.Idle:
		// 空闲状态，尝试激活连接
		if hc.manager.conn != nil {
			hc.manager.conn.Connect()
			if hc.debug {
				log.Printf("[INFO] HealthChecker.checkHealth: 激活空闲连接")
			}
		}
	case connectivity.Ready:
		// 连接正常，无需操作
	case connectivity.Connecting:
		// 正在连接中，等待
		if hc.debug {
			log.Printf("[INFO] HealthChecker.checkHealth: 连接正在建立中")
		}
	}
}

// attemptReconnect 尝试重新连接
func (hc *HealthChecker) attemptReconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := hc.manager.Reconnect(ctx, ""); err != nil {
		log.Printf("[ERROR] HealthChecker.attemptReconnect: 重连失败: %v", err)
	} else {
		if hc.debug {
			log.Printf("[INFO] HealthChecker.attemptReconnect: 重连成功")
		}
	}
}

// IsRunning 检查健康检查器是否正在运行
func (hc *HealthChecker) IsRunning() bool {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return hc.running
}