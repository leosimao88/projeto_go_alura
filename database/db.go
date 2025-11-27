package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é a variável global para o objeto de conexão GORM
var DB *gorm.DB

func ConectaComBancoDeDados() {
	// Obtendo variáveis de ambiente
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")

	// Usando fmt.Sprintf para construir a string de conexão de forma segura e limpa
	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		endereco, usuario, senha, nomeBanco, portaBanco,
	)

	// Conectando com o banco de dados.
	var errConexao error
	// Adicionando um objeto vazio de Configuração do GORM (boa prática)
	DB, errConexao = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{}) 

	if errConexao != nil {
		// Pânico (Panic) com o erro real para facilitar a depuração.
		log.Panicf("Erro ao conectar com banco de dados: %v", errConexao)
	}

	// Rodando AutoMigrate e verificando o erro (não ignorando)
	errMigrate := DB.AutoMigrate(&models.Aluno{})
	if errMigrate != nil {
		log.Panicf("Erro ao rodar AutoMigrate: %v", errMigrate)
	}
	
	log.Println("Conexão com o banco de dados e AutoMigrate realizados com sucesso!")
}