This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.

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