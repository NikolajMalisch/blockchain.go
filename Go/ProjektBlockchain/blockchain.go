package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block einen benutzerdefinierten Blocktyp, um die Daten unserer Blockchain zu speichern
type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

// Blockchain einen benutzerdefinierten Typ, der unsere Blöcke enthält
type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

// CalculateHash Block-Hash für unsere Blockchain ableiten, indem wir den vorherigen Block-Hash,
//die aktuellen Blockdaten, den Zeitstempel und den PoW mit demSHA256-Algorithmushashen
func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// Mine Erstellen wir eine Methode für unseren Typ, die den PoW-Wert erhöht und
//den Block hash berechnet, bis wir einen gültigen Hash erhalten
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

// CreateBlockchain Funktion, die einen Genesis-Block für unsere Blockchain erstellt und eine neue Instanz des Typs zurückgibt
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	//Hash-Wert unseres Genesis-Blocks auf einen Wert setzen. Da dies der erste Block in der Blockchain ist,
	//gibt es keinen Wert für den vorherigen Hash und die Dateneigenschaft ist leer. 0
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// AddBlock eineMethode für einen Typ erstellt haben, die Folgendes tut
//erfasst Transaktionsdetails (Absender, Empfänger und Überweisungsbetrag)
//erzeugt einen neuen Block mit Transaktionsdetails
//erstellt einen neuen Block unter Verwendung des Hash werts des vorherigen Blocks,
//der aktuellen Blockdaten und des generierten PoW fügt den neu erstellten Block zur Blockchain hinzu
func (b *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.Mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

// IsValid berechnet den Hash-Wert jedes Blocks neu, vergleicht ihn mit den gespeicherten Hash-Werten der anderen Blöcke
//und prüft, ob der vorherige Hash-Wert eines anderen Blocks gleich dem Hash-Wert des vorherigen Blocks ist.
//Wenn eine der beiden Prüfungen fehlschlägt, wurde die Blockchain manipuliert.
func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.AddBlock("Alice", "Bob", 5)
	blockchain.AddBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.IsValid())
}
