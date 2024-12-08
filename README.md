# Informatik 1

Dieses Repo enthält Material zur Vorlesung Informatik 1
im Kurs TEIT24A an der DHBW Mannheim.

## Software

In der Vorlesung wird die Programmiersprache [Go](https://go.dev) verwendet.
Dies ist eine relativ junge, freie Programmiersprache, die dank einer einfachen
Entwicklungs-Toolchain und einer guten Standardbibliothek einen leichten Einstieg
und kurze Entwicklungszyklen ermöglicht.
Gleichzeitig ist Go hochgradig praxisrelevant:
Die Sprache wird u.A. im Web- bzw. Cloud-Umfeld oder z.B. für service-orientierte Softwaresysteme verwendet.

Hier eine Zusammenfassung der verwendeten/empfohlenen Software:

| Software | Download-Link | Erläuterung
| --- | --- | ---
| Go-[Compiler](https://de.wikipedia.org/wiki/Compiler) | <https://go.dev>                     | Das wichtigste Werkzeug beim Programmieren: Übersetzt den [Quellcode](https://de.wikipedia.org/wiki/Quelltext) in eine [ausführbare Datei](https://de.wikipedia.org/wiki/Ausf%C3%BChrbare_Datei).
| Visual Studio Code                                    | <https://code.visualstudio.com>      | [Code-Editor](https://de.wikipedia.org/wiki/Editor_(Software)) bzw. [Entwicklungsumgebung](https://de.wikipedia.org/wiki/Integrierte_Entwicklungsumgebung) zum Schreiben von Quelltext.
| GoLand                                                | <https://www.jetbrains.com/de-de/go> | Eine Alternative zu VsCode. Gut integriert und einfach einzusetzen, "batteries included". Für Studierende kostenlos.
| Git                                                   | <https://git-scm.com>                | Eine [Versionsverwaltung](https://de.wikipedia.org/wiki/Versionsverwaltung) für Quellcode.
| Go-Playground                                         | <https://go.dev/play>                | Ein einfacher Online-Editor und -Compiler für Go-Code. Gut zum schnellen Ausprobieren.
| Replit                                                | <https://replit.com>                 | Eine komplexere Online-Entwicklungsumgebung für diverse Programmiersprachen.

Dies alles sollte auf jedem Windows- oder Linux-PC und auch auf jedem Mac
ohne Admin-Rechte installierbar sein.
Falls es dabei dennoch Probleme gibt, ist auch die Verwendung eines webbasierten
Programmiersystems denkbar (z.B. [Go Playground](https://go.dev/play) oder [Replit](https://replit.com), s.o.).

## Bemerkungen zur Software-Verwendung und -Installation

### VsCode

In der Vorlesung und Prüfung wird VsCode als Editor verwendet. Wer möglichst ähnliche Programmierumgebung bei sich einrichten will,
sollte sich also den Go-Compiler von <https://go.dev> sowie VsCode installieren.
*Wichtig dabei*: Sobald zum ersten Mal eine Go-Quelldatei mit VsCode geöffnet wird, wird zuerst die Installation einer *Go-Erweiterung*
und anschließend noch die Installation diverser Tools angeboten. Dies sollte alles angenommen werden.
VsCode installiert aber nicht selbständig den Go-Compiler, dieser muss vorab selbständig installiert werden.

### GoLand

Gegen die Verwendung von GoLand ist grundsätzlich nichts einzuwenden, dies ist ebenfalls eine sehr gute Entwicklungsumgebung.
Die Installation ist etwas einfacher, die Umgebung ist stärker automatisiert. Die Bedienung ist vom Konzept her ähnlich wie bei VsCode,
unterscheidet sich aber in Details.
Als Vorbereitung für die Prüfung ist es daher ratsam, VsCode wenigstens einmal ausprobiert zu haben.

### Browser-Tools

Die webbasierten Umgebungen können sehr nützlich sein, um schnell mal Beispiele auszuprobieren oder mitzucoden,
wenn man keinen vollwertigen Rechner hat (z.B. Tablet oder stark eingeschränkten Firmenlaptop).
Insbesondere Replit kann auch nützlich sein, um zu mehreren online zusammenzuarbeiten.
Ein vollwertiger Ersatz für eine eigene, lokal installierte Umgebung sind diese aber nicht.

Falls es gar nicht gelingt, einen eigenen Rechner einzurichten, können auch Tools wie [GitHub Codespaces](https://github.com/features/codespaces)
oder [GitPod](https://www.gitpod.io) nützlich sein. Diese basieren auf der Web-Version von VsCode und bieten virtualisierte Entwicklungsumgebungen
an, die einen eigenen Rechner simulieren. Um hiermit effektiv zu arbeiten, sollte dann aber auch Git verwendet werden.

### Git

[Git](https://git-scm.com) ist eine Versionsverwaltung für Quellcode, deren Hauptzweck es ist, gemeinsam  an Quellcode zu arbeiten und sich dabei
möglichst nicht gegenseitig mit Änderungen zu stören.
Services wie [GitHub](https://github.com) setzen darauf auf und bieten einen zentralen Speicherort für Quellcode.
Dadurch ist Git mittlerweile zum weltweiten Quasi-Standard für die Verwaltung von Software-Quellcode geworden.

Sämtlicher Code für diese Vorlesung wird über GitHub bereitgestellt.
Die Verwendung von Git ist dabei nicht zwingend notwendig und nicht Kernbestandteil der Vorlesung.
Sich mit den einfachsten Git-Kommandos und Arbeitsweisen vertraut zu machen, vereinfacht aber auch die Arbeit mit dem Code aus der Vorlesung.
Beispiele dafür werden in der Vorlesung besprochen.

Neben dieser allgemeinen Anwendung wird Git auch von Go selbst als Werkzeug für die Verwaltung und Bereitstellung von
Paketen genutzt. Ob dies in der Vorlesung von Relevanz sein wird, ist jetzt am Anfang noch nicht klar.
Zu einer funktionierenden und vollständigen Go-Installation gehört Git dadurch aber auf jeden Fall dazu und sollte bei Ihnen installiert sein.
