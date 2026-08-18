# =============================================================================
# STAGE 1: Builder - Compila Vue con Vite
# =============================================================================
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .

RUN npm run build

# =============================================================================
# STAGE 2: Runner - Servidor Nginx
# =============================================================================
FROM nginx:1.25-alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8081

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8081/health.html || exit 1

CMD ["nginx", "-g", "daemon off;"]
