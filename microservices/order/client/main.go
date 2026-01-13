package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	orderpb "github.com/Mellanie-Marques/microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Conectar ao servidor gRPC
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	client := orderpb.NewOrderClient(conn)

	// Itens de exemplo
	items := []*orderpb.OrderItem{
		{ProductCode: "prod1", UnitPrice: 10.0, Quantity: 2},
		{ProductCode: "prod2", UnitPrice: 5.0, Quantity: 1},
	}

	// Calcular total
	var total float32
	for _, it := range items {
		total += it.UnitPrice * float32(it.Quantity)
	}

	// Exibir visualização da requisição
	fmt.Println("========================================")
	fmt.Println("  Cliente de teste — Criar novo pedido")
	fmt.Println("========================================")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Produto\tPreço\tQuantidade\tSubTotal")
	for _, it := range items {
		fmt.Fprintf(w, "%s\t%.2f\t%d\t%.2f\n", it.ProductCode, it.UnitPrice, it.Quantity, it.UnitPrice*float32(it.Quantity))
	}
	w.Flush()
	fmt.Printf("----------------------------------------\nTotal: %.2f\n", total)

	// Pausa para o usuário confirmar
	fmt.Print("Pressione Enter para enviar o pedido... ")
	bufio.NewReader(os.Stdin).ReadString('\n')

	// Montar requisição
	req := &orderpb.CreateOrderRequest{
		CostumerId: 123,
		OrderItems: items,
		TotalPrice: total,
	}

	// Definir timeout para a requisição
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Chamar o método Create
	resp, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("Erro ao criar pedido: %v", err)
	}

	// Exibir resposta de forma visual
	fmt.Println("\n========================================")
	fmt.Println(" Resultado da criação do pedido ")
	fmt.Println("========================================")
	fmt.Printf("ID do Pedido: %d\n", resp.OrderId)
	fmt.Printf("Enviado em: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("========================================")
}
