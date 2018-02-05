package helper

import (
    "crypto/sha256"
    "fmt"
    "github.com/google/uuid"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/cloudflare/cfssl/log"
    "encoding/hex"
    "strings"
)

const DIFFICULTY = 4

func GetHash(text string) string {
    algorithm := sha256.New()
    algorithm.Write([]byte(text))
    return hex.EncodeToString(algorithm.Sum(nil))
}

func IsHashCorrectDifficulty(hash string) bool {
    difficulty := strings.Repeat("0", DIFFICULTY)
    return hash[:DIFFICULTY] == difficulty
}

func ValidProof(lastProof int, proof int, lastHash string) bool {
    sprintf := fmt.Sprintf("%d%d%s", lastProof, proof, lastHash)
    hash := GetHash(sprintf)
    log.Info(fmt.Sprintf(`Hash for "%d%d%s" results in %s`, lastProof, proof, lastHash, hash))
    return IsHashCorrectDifficulty(hash)
}

func GetUrl(url string) map[string]interface{} {
    res, _ := http.Get(url)
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil
    }
    var response map[string]interface{}

    body, _ := ioutil.ReadAll(res.Body)
    json.Unmarshal(body, response)
    return response
}

func ConvertInterface(input interface{}, output interface{}) {
    result, _ := json.Marshal(input)
    json.Unmarshal(result, &output)
}

func GenerateAddress() string {
    return uuid.New().String()
}
