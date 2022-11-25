
# Blockchain.go

 Die Eingegebenen Daten sollen unveränderbar abgespeichert werden. Hierfür eignet sich eine eigene Blockchain. Der Kunde erwartet die Implementierung in der Programmiersprache Golang.

## Documentation

[Documentation](https://github.com/NikolajMalisch/blockchain.go.git)

## Installation

git clone https://github.com/NikolajMalisch/blockchain.go.git 

cd blockchain.go

import ()

    // Block einen benutzerdefinierten Blocktyp, um die Daten unserer Blockchain zu speichern
type Block struct {

    // Blockchain einen benutzerdefinierten Typ, der unsere Blöcke enthält
type Blockchain struct {

    // CalculateHash Block-Hash für unsere Blockchain ableiten, indem wir den vorherigen Block-Hash,
    //die aktuellen Blockdaten, den Zeitstempel und den PoW mit demSHA256-Algorithmushashen
func (b Block) CalculateHash() string {

    // Mine Erstellen wir eine Methode für unseren Typ, die den PoW-Wert erhöht und
    //den Block hash berechnet, bis wir einen gültigen Hash erhalten
func (b *Block) Mine(difficulty int) {

    // CreateBlockchain Funktion, die einen Genesis-Block für unsere Blockchain erstellt und eine neue Instanz des Typs zurückgibt
func CreateBlockchain(difficulty int) Blockchain {

    //Hash-Wert unseres Genesis-Blocks auf einen Wert setzen. Da dies der erste Block in der Blockchain ist,
    //gibt es keinen Wert für den vorherigen Hash und die Dateneigenschaft ist leer. 0
return Blockchain{

    // AddBlock eineMethode für einen Typ erstellt haben, die Folgendes tut
    //erfasst Transaktionsdetails (Absender, Empfänger und Überweisungsbetrag)
    //erzeugt einen neuen Block mit Transaktionsdetails
    //erstellt einen neuen Block unter Verwendung des Hash werts des vorherigen Blocks,
    //der aktuellen Blockdaten und des generierten PoW fügt den neu erstellten Block zur Blockchain hinzu
func (b *Blockchain) AddBlock(from, to string, amount float64) {
    // IsValid berechnet den Hash-Wert jedes Blocks neu, vergleicht ihn mit den gespeicherten Hash-Werten der anderen Blöcke
    //und prüft, ob der vorherige Hash-Wert eines anderen Blocks gleich dem Hash-Wert des vorherigen Blocks ist.
    //Wenn eine der beiden Prüfungen fehlschlägt, wurde die Blockchain manipuliert.
func (b Blockchain) IsValid() bool {

func main() {
kchain := CreateBlockchain(2)

// record transactions on the blockchain for Alice, Bob, and John
blockchain.AddBlock("Alice", "Bob", 5)
blockchain.AddBlock("John", "Bob", 2)

// check if the blockchain is valid; expecting true
fmt.Println(blockchain.IsValid())

go run blockchain.go


Quelle für bericht // https://medium.com/@pkthakur01/why-golang-for-blockchain-40f874f2ce1b