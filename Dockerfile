FROM ghcr.io/typst/typst:latest AS compile

WORKDIR /usr/local/app
COPY resume.typ ./
COPY resume.yaml ./
RUN typst compile resume.typ resume.pdf

FROM nginx:latest
WORKDIR /usr/share/nginx/html

COPY <<'EOF' /etc/nginx/conf.d/default.conf
server {
    listen 80;
    location / {
        root /usr/share/nginx/html;
        index /resume.pdf;
    }
}
EOF

RUN cat /etc/nginx/conf.d/default.conf

COPY --from=compile /usr/local/app/resume.pdf /usr/share/nginx/html/resume.pdf
