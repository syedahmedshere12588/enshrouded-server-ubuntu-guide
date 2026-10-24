package main
import ("encoding/json";"fmt";"os")
func main() {
    f, _ := os.Open("config.json"); defer f.Close()
    var cfg map[string]interface{}; json.NewDecoder(f).Decode(&cfg)
    fmt.Printf("Loaded %d items\n", len(cfg))
}
