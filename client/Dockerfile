# ---- Build Next.js Frontend ----
FROM public.ecr.aws/docker/library/node:20-alpine AS frontend-builder

WORKDIR /app

COPY client/package*.json ./
RUN npm install

COPY client/ ./
RUN npm run build

# ---- Serve with Next.js built-in server ----
FROM node:20-alpine

WORKDIR /app

COPY --from=frontend-builder /app ./

EXPOSE 3000

CMD ["npm", "start"]

