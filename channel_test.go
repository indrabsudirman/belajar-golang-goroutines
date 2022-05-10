package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Indra Bayu"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Indra Sudirman"

}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) { //Hanya mengirim Channel
	time.Sleep(2 * time.Second)
	channel <- "Indra Sudirman"
}

func OnlyOut(channel <-chan string) { //Hanya menerima Channel
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Indra"
	channel <- "Bayu"
	channel <- "Sudirman"

	fmt.Println(len(channel)) //melihat jumlah data di buffered hasil = 3, karena belum diambil
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(cap(channel)) //melihat panjang buffered
	//Kalo di windows, error dead lock channel tidak tampil. Tapi hanya stuck cursor diem saja
	//seperti lagi nunggu sesuatu. Padahal mungkin channel belum diambil atau belum diisi.
	//Makanya menggunakan buffered

	//Bisa juga buffered channel didalam Goroutines
	//Buat anonymous func, seperti ini
	fmt.Println("Channel dalam Goroutines")
	go func() {
		channel <- "Indra"
		channel <- "Bayu"
		channel <- "Sudirman"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Selesai")

}

//Penggunaan Range di Channel untuk kasus jika tidak jelas berapa banyak yang akan dikirim
//ke channel. Jadi bisa menggunakan range (looping range)
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		//Jangan lupa harus close channel, jika tidak akan error dead lock
		// di windows error dead locknya, terlihat kursor stuck
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//Untuk mengambil dari beberaoa channel
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

//Untuk mengambil jika data dichannel kosong, maka bisa pakai default(pesannya)
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//Untuk mengambil dari beberaoa channel
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}

}
