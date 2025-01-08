####################################################
# GOLANG BUILDER
####################################################
FROM golang:1.20 AS go_builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers sources locaux
COPY . /app

RUN go mod edit -replace github.com/Sirupsen/logrus=github.com/sirupsen/logrus@v1.9.3 && \
    go mod tidy

# Construire le binaire

RUN go build -ldflags "-s -w -X main.Version=$(cat VERSION || echo 'v1.0.0') -X main.BuildTime=$(date -u +%Y%m%d)" -o /bin/avscan

####################################################
# PLUGIN BUILDER
####################################################
FROM ubuntu:22.04 AS plugin_builder

# Créer un utilisateur et un groupe
RUN groupadd -r malice && \
    useradd --no-log-init -r -g malice malice && \
    mkdir /malware && \
    chown -R malice:malice /malware

# Copier les fichiers nécessaires
COPY installer.tar /tmp/installer.tar

# Installer Bitdefender et les dépendances
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates psmisc && \
    tar -xvf /tmp/installer.tar -C /opt/ && \
    chmod -R +x /opt/installer && \
    bash /opt/installer/installer.sh && \
    mkdir -p /opt/BitDefender-scanner && \
    cp -r /opt/installer/* /opt/BitDefender-scanner/ && \
    /opt/BitDefender-scanner/bin/bdscan --update || true && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
    chmod -R +x bd_fix.sh && \
    bash bd_fix.sh


# Copier le binaire construit dans la première étape
COPY --from=go_builder /bin/avscan /bin/avscan

# Assurer les permissions
WORKDIR /malware
RUN chown -R malice:malice /malware

# Définitions d'entrée et commande par défaut
USER malice
ENTRYPOINT ["/bin/avscan"]
CMD ["--help"]

