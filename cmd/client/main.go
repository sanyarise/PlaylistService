package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/caarlos0/env"
	playlist "github.com/sanyarise/playlist/internal/delivery/grpc/interface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Config struct {
	Host string `env:"SRV_HOST" envDefault:"localhost"`
	Port string `env:"SRV_PORT" envDefault:"8000"`
}

var (
	config Config
	once   sync.Once
)

func NewConfig() *Config {
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			log.Fatalf("Can't load configuration: %s", err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Load config successful %v", string(configBytes))
	})
	return &config
}

func CreateSong(c playlist.PlaylistClient) {
	var title string
	var duration uint32

	fmt.Println("enter title of song:")
	fmt.Scan(&title)
	fmt.Println("enter duration of song:")
	fmt.Scan(&duration)
	req := &playlist.CreateSongRequest{
		Title:    title,
		Duration: duration,
	}
	res, err := c.CreateSong(context.Background(), req)
	if err != nil {
		log.Println("cannot create song: ", err)
		return
	}
	log.Printf("song created success with id: %s", res.Id)
}

func GetSong(c playlist.PlaylistClient) {
	var id string
	fmt.Println("Enter id of song:")
	fmt.Scan(&id)
	req := &playlist.GetSongRequest{
		Id: id,
	}
	res, err := c.GetSong(context.Background(), req)
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't get song. Status: %v, err: %v", status, err)
		return
	}
	log.Printf("get song success: song: %v", res.Song)
}

func UpdateSong(c playlist.PlaylistClient) {
	var id, title string
	var duration uint32
	fmt.Println("Enter id: ")
	fmt.Scan(&id)
	fmt.Println("Enter title: ")
	fmt.Scan(&title)
	fmt.Println("Enter duration: ")
	fmt.Scan(&duration)
	req := &playlist.UpdateSongRequest{
		Id:       id,
		Title:    title,
		Duration: duration,
	}
	res, err := c.UpdateSong(context.Background(), req)
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't update song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("update song success", res)
}

func DeleteSong(c playlist.PlaylistClient) {
	var id string
	fmt.Println("Enter song id: ")
	fmt.Scan(&id)
	req := &playlist.DeleteSongRequest{
		Id: id,
	}
	res, err := c.DeleteSong(context.Background(), req)
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't delete song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("delete song success", res)
}

func AddSong(c playlist.PlaylistClient) {
	var id string
	fmt.Println("Enter id: ")
	fmt.Scan(&id)
	req := &playlist.AddSongRequest{
		Id: id,
	}
	res, err := c.AddSong(context.Background(), req)
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't add song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("add song success", res)
}

func PlaySong(c playlist.PlaylistClient) {
	res, err := c.PlaySong(context.Background(), &playlist.PlaySongRequest{})
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't play song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("play song success", res)
}

func PauseSong(c playlist.PlaylistClient) {
	res, err := c.PauseSong(context.Background(), &playlist.PauseSongRequest{})
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't pause song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("pause song success", res)
}

func NextSong(c playlist.PlaylistClient) {
	res, err := c.NextSong(context.Background(), &playlist.NextSongRequest{})
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't next song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("next song success", res)
}

func PrevSong(c playlist.PlaylistClient) {
	res, err := c.PrevSong(context.Background(), &playlist.PrevSongRequest{})
	if err != nil {
		status, _ := status.FromError(err)
		log.Printf("can't prev song. Status: %v, err: %v", status, err)
		return
	}
	log.Println("prev song success", res)
}

func main() {
	cfg := NewConfig()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't dial server: ", err)
	}
	log.Println("dial server success")

	client := playlist.NewPlaylistClient(conn)
	for {
		fmt.Println("Input type of operation (create, get, update, delete, add, play, pause, next, prev, exit): ")
		var op string
		fmt.Scan(&op)
		s := strings.ToLower(op)
		switch {
		case s == "create":
			CreateSong(client)
		case s == "get":
			GetSong(client)
		case s == "update":
			UpdateSong(client)
		case s == "delete":
			DeleteSong(client)
		case s == "add":
			AddSong(client)
		case s == "play":
			PlaySong(client)
		case s == "pause":
			PauseSong(client)
		case s == "next":
			NextSong(client)
		case s == "prev":
			PrevSong(client)
		case s == "exit":
			return
		default:
			fmt.Println("unknown operation")
		}
	}
}
