# factro-task-replacer

## Was macht das Programm?
Das Programm ist in der Lage, mehrere "Tasks" in Factro gleichzeitig zu aktualisieren. Es ist in der Lage, alle Tasks für einen bestimmten Benutzer herunterzuladen, ein bestimmtes Feld, z.B. den Titel, nach einem bestimmten Wert, z.B. "04.2020", zu durchsuchen und diesen Wert überall durch einen anderen Wert , z.B. "04.2021" zu ersetzen. Anschließend werden die aktualisierten Tasks wieder zum Factro Server geschickt, der sie in der Cloud aktualisiert. 

## Wie benutze ich das Programm?
1. factro-task-replacer.exe herunterladen und in ein Verzeichnis speichern.
2. Im selben Verzeichnis einen Ordner mit dem Namen "config" erstellen.
3. Im Ordner "config" eine Datei mit dem Namen "api_user_token_.txt" erstellen.
4. In diese Textdatei den API-Token (JWT) des Nutzers kopieren, für den die Tasks aktualisiert werden sollen. Dabei sollte ein Testnutzer verwendet werden, dem nur das Projekt zugeordnet ist, dessen Tasks aktualisiert werden sollen. Das API Token lässt sich in Factro generieren.
5. Die Textdatei abspeichern. 
6. "factro-task-replacer.exe" durch Doppelklicken ausführen.
7. Die Anweisungen auf dem Bildschirm befolgen.
8. Eingaben durch die Eingabetaste (Enter) bestätigen.
9. Nach erfolgreicher Ausführung in der Webanwendung überprüfen, ob die Aktualisierung erfolgreich durchgeführt werden konnte.


## Troubleshooting
| Problem                                                                         | Mögliche Ursachen                                                               |
|---------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| Es gibt keine Fehlermeldung, die Tasks werden aber trotzdem nicht aktualisiert. | Das angegebene Administrator-Token besitzt nicht die notwendigen Berechtigungen |
