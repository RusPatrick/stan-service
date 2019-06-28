package main

func main() {
	stop := make(chan struct{})
	go services.StartStanService(stop)
	<-stop
}
