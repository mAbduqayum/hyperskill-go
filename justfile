default:
    @just --list --unsorted

mk start end:
    #!/usr/bin/env bash
    range="{{start}}_{{end}}"
    for i in $(seq {{start}} {{end}}); do
        echo "==> Creating structure for $i"
        mkdir -p $range/${i}_problem/tests
        touch $range/${i}_problem/${i}_problem{,_test}.go
    done
