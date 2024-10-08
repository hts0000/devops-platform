# Plan
## step1——应用部署记录平台
- [ ] 侧边栏
- [ ] 创建业务
- [ ] 删除业务
- [ ] 修改业务
- [ ] 查看业务
- [ ] 查看业务列表
- [ ] 创建项目
- [ ] 删除项目
- [ ] 修改项目
- [ ] 查看项目
- [ ] 查看项目列表
- [ ] 创建应用
- [ ] 删除应用
- [ ] 修改应用
- [ ] 查看应用
- [ ] 查看应用列表

### 业务
大的业务线，如海外供应链、门店人事系统等

业务归属：
- 业务负责人：业务leader

### 项目
隶属于业务线下的项目，如门店人事系统中的MINISO门店人事、TOPTOY门店认识等。

项目还区分测试环境、预发环境和生产环境，使用TEST、PRE、PRD进行区分。

项目归属：
- 产品负责人：负责项目需求提出和拆分
- 开发负责人
- 测试负责人
- DBA
- SRE

### 应用
隶属于项目下的前后端应用，如MINISO门店人事项目中的storehr-core应用。

应用归属：
- 开发负责人
- 开发人员：参与该应用的开发
- 测试负责人：
- 测试人员：参与该应用的测试

应用名规范: `<miniso/toptoy/wow>-<appname>-<web/java/go/php>-<test/pre/prd>`

应用版本规范：
- 测试和预发环境：`<test/pre>-<timestamp:buildnum>`
- 生产环境：`<prd>-<tag:buildnum>`

### UI设计
#### 创建业务页面

#### 创建项目页面

#### 创建应用页面
![](https://cdn.jsdelivr.net/gh/hts0000/images/202410071153949.png)

## step2——自动化应用部署
- [ ] 创建应用自动创建Jenkins条目(需根据业务、项目和环境创建文件夹)
- [ ] Jenkins条目创建后自动触发构建
- [ ] 构建成功后自动关联SLS日志库
- [ ] 可选是否接入ARMS
- [ ] 展示应用Gitlab、Jenkins、K8S、SLS等链接

## step3——流程化应用发布
- [ ] 定义发版日
- [ ] 定义发版时间表
- [ ] 定义开发需求
- [ ] 定义开审批人、发负责人、产品、测试、DBA、SRE等角色
- [ ] 定义角色权限
- [ ] 创建生产环境发布任务，关联需求、变更的应用、变更的配置、审批人、业务负责人、项目负责人、开发负责人、产品、测试、DBA、SRE，记录应用发布关联的需求、应用、配置、审批人、版本等信息

### 发版
- 常规发版：只能在定义的发版日及时间表中发版
- 特殊发版：任意日期和时间段发版

### 角色
- 审批人
- 业务负责人
- 项目负责人
- 开发负责人
- 测试负责人
- 产品负责人
- DBA负责人
- SRE负责人

### 发版任务状态流转
1. 产品负责人对其拥有权限的项目创建常规发版或特殊发版任务，任务关联需求。创建成功后任务流转至开发负责人
2. 开发负责人添加该任务关联的应用，提供应用关联数据库、配置变更、SQL变更、发布tag、应急回滚等信息。提供后任务流转至测试负责人
3. 测试负责人确认需求及关联的应用已在预发环境验收，提供验收报告等信息。验收通过后任务流转至审批人，抄送业务负责人
4. 审批人审批发版任务。审批通过后任务流转至DBA(如有SQL)
5. DBA执行SQL变更，提供应急回滚方案。完成后任务流转至SRE
6. SRE确认变更影响，提供应急回滚方案。完成后任务归档，通知该发版任务所有相关人员

## step4——添加应用告警功能
- [ ] 宏观维度的——链路追踪，从dns->waf->lb->app的链路全景
- [ ] 微观维度的——应用业务数据、基础设施(dns/waf/lb/mq/db/ecs/k8s)数据
- [ ] 规范应用数据上报
- [ ] 创建应用告警规则
- [ ] 告警通知(短信、邮箱、电话、飞书、企微等)
- [ ] 告警静默
- [ ] 告警认领
- [ ] 告警恢复

### 应用数据上报
- ARMS+云监控2.0
- exporter+prometheus

# How to use
## Connect to MySQL
```bash
# 初始化
# -e MYSQL_ROOT_PASSWORD=123456: 注入环境变量，配置root用户密码
# -e MYSQL_DATABASE=devops: 注入环境变量，自动创建devops库
# -v /my/custom:/etc/mysql/conf.d: 将本地配置文件挂载到docker中
# -v /my/own/datadir:/var/lib/mysql: 使用本地目录作为数据库存储目录
# -v /my/initsql:/docker-entrypoint-initdb.d: 挂载到/docker-entrypoint-initdb.d目录的脚本，容器初始化时会自动执行
# -p 13306:3306: 将容器3306端口映射到本机13306端口
# --character-set-server=utf8mb4: 配置数据库字符集
# --collation-server=utf8mb4_unicode_ci: 配置数据表字符集
# --restart on-failure:3: 容器异常退出时，总是重启容器，最多重启3次
docker run --name devops-mysql \
            --restart on-failure:3 \
            -v /my/custom:/etc/mysql/conf.d \
            -v /my/own/datadir:/var/lib/mysql \
            -v /somepath/devops/mysql/sql:/docker-entrypoint-initdb.d \
            -v /somepath/devops/mysql/sql:/docker-entrypoint-initdb.d \
            -e MYSQL_ROOT_PASSWORD=123456 \
            -e MYSQL_DATABASE=devops \
            -p 13306:3306 \
            -d mysql:8.3.0 \
            --character-set-server=utf8mb4 \
            --collation-server=utf8mb4_unicode_ci

# 测试使用
docker run --name devops-mysql-test \
            --restart on-failure:3 \
            -e MYSQL_ROOT_PASSWORD=123456 \
            -e MYSQL_DATABASE=devops \
            -p 13306:3306 \
            -d mysql:8.3.0 \
            --character-set-server=utf8mb4 \
            --collation-server=utf8mb4_unicode_ci

# 连接
docker exec -it devops-mysql mysql -uroot -p123456

# migrate
npx prisma migrate dev --name init

# update schema
npx prisma db push
```