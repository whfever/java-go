
# 并发
## 问题根源
CPU缓存  分时  重排序
## 控制手段
阻塞（锁）  信号量
非阻塞同步 
无同步方案 线程隔离

## Java 解决机制
volatile retreenlock synchronized 

## 线程

1. 线程状态
2. 线程同步Synchronized RetreeLock

3. 线程协作join  wait/await notify notifyAll

## 锁

乐观CAS/悲观
自旋锁 避免线程切换开销/
公平锁 非公平锁
可重入锁/非可重入锁

## JUC 工具类
集合 线程池 工具类
 CountDownLatch,CyclicBarrier,Semaphore
ReentrantLock,Synchronized,ReentrantReadWriteLock,CopyOnWriteArrayList,

# JVM
## JMM

## 类加载

## 内存结构

## 垃圾回收

## 调优排错