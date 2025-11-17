# CHECKUSER 🕵️‍♂️

## Autor
Glemison C. DuTra ([DuTra](https://t.me/DTunnel))

## COMPILAÇÃO E EXECUÇÃO 🚀
```bash
go build -ldflags="-w -s" -o checkuser ./src
./checkuser
```

## INICIAR CHECKUSER 🚀
```bash
./checkuser --start --port 8080
```

## INSTALAÇÃO NO SEU SERVIDOR 🛠️
### 1. Usando o repositório oficial da JJ Secure VPN
```bash
sudo apt update && sudo apt install -y git
git clone https://github.com/jjsecurevpn-vpn/CheckUser-Secure.git
cd CheckUser-Secure

# Instale o Go 1.21+ (se ainda não tiver)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=/usr/local/go/bin:$PATH

go build -ldflags="-w -s" -o checkuser ./src
./checkuser --start --port 8080
```

### 2. Instalador automático legado (repositório original)
```bash
bash <(curl -sL https://n9.cl/yo2nc)
```
> **Nota:** o script legado baixa o projeto original de Glemison DuTra. Use apenas se você optar por continuar com a versão antiga.
