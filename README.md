
Blockchain.go
Dieses Projekt implementiert eine einfache Blockchain in Golang, die Eingabedaten unveränderbar speichert.

Dokumentation
Die Dokumentation ist hier verfügbar.

Installation
Klonen Sie das Repository mit git clone https://github.com/NikolajMalisch/blockchain.go.git
Wechseln Sie in das Verzeichnis cd blockchain.go
Führen Sie go run blockchain.go aus, um die Anwendung zu starten.
Verwendung
Block
Ein Block ist ein benutzerdefinierter Typ, der die Daten unserer Blockchain speichert. Ein Block enthält folgende Felder:

PrevHash: der Hash-Wert des vorherigen Blocks in der Blockchain
Data: die Daten, die im Block gespeichert werden sollen
Timestamp: ein Zeitstempel, der angibt, wann der Block erstellt wurde
Nonce: ein Wert, der für den Proof-of-Work-Algorithmus verwendet wird, um einen gültigen Hash-Wert zu finden
Hash: der Hash-Wert des aktuellen Blocks
Blockchain
Eine Blockchain ist ein benutzerdefinierter Typ, der eine Sammlung von Blöcken enthält. Eine Blockchain enthält folgende Felder:

Blocks: eine Liste von Blöcken, die die Blockchain bilden
Die Blockchain unterstützt folgende Methoden:

AddBlock(from, to string, amount float64): fügt einen neuen Block zur Blockchain hinzu, der Transaktionsdetails enthält
IsValid() bool: prüft, ob die Blockchain manipuliert wurde, indem sie den Hash-Wert jedes Blocks neu berechnet und überprüft, ob der vorherige Hash-Wert eines anderen Blocks gleich dem Hash-Wert des vorherigen Blocks ist.
Beispiel
Ein Beispiel für die Verwendung der Blockchain ist in der main-Funktion enthalten:

go
Copy code
func main() {
    kchain := CreateBlockchain(2)

    // Transaktionen auf der Blockchain für Alice, Bob und John aufzeichnen
    blockchain.AddBlock("Alice", "Bob", 5)
    blockchain.AddBlock("John", "Bob", 2)

    // Überprüfen, ob die Blockchain gültig ist; erwartet wird true
    fmt.Println(blockchain.IsValid())
}
Contributing
Wenn Sie einen Fehler gefunden haben oder einen Vorschlag zur Verbesserung des Codes oder der Dokumentation haben, können Sie gerne einen Pull-Request erstellen oder ein Issue eröffnen.

Lizenz
Dieses Projekt ist unter der MIT-Lizenz lizenziert. Weitere Informationen finden Sie in der LICENSE-Datei.
