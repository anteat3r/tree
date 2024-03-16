Realtime generace fraktálního stromu.
Po spuštěné zmáčkněte `O` pro načtení základní konfigurace.
Nejjednodušší mi přijde prostě vyzkoušet, co jednotlivé klávesy dělají.
Export pro Windows jsem netestoval, ale měl by fungovat.

P.S.
Ze třech stromů lze sestavit (limitně) přesnou Kochovu vločku.
Můžete na odpovídající konfigraci zkusit přijít sama nebo do
`config.json` vložit konfiguraci z `koch.json`.

### Ovládání
Dvojice kláves vždy snižují/zvyšují jednu hodnotu.
Klávesy jsou kalibrovány na anglickou klávesnici.

`U` + `L` -> úhel větví
`K` + `J` -> faktor zmenšení další iterace 
`,` + `.` -> poměr délky hlavní a 2 vedlejších větví
`[` + `]` -> délka hlavní větve první iterace
`9` + `0` -> horizontální offset
`1` + `2` -> vertikální offset
`Q` + `W` -> vzdálenost od středu kde stromy začínají
`-` + `=` -> počet renderovaných iterací
`M` + `N` -> počet renderovaných stromů
`O` -> načíst konfiguraci z `config.json`
`P` -> uložit momentální konfiguraci do `config.json`
