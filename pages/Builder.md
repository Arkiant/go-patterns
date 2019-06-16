# Builder

## Descripción

El patrón Builder nos ayuda a construir objetos complejos sin instanciar directamente su estructura, ni escribir la lógica que requieren. Imagina un objeto que podría tener docenas de campos que son estructuras más complejas. Ahora imagina que tienes muchos objetos con estas características, y podrías tener más. No queremos escribir la lógica para crear todos estos objetos en el paquete que sólo necesita usar los objetos.

Por ejemplo, usted usará casi la misma técnica para construir un coche que para construir un autobús, excepto que serán de diferentes tamaños y número de plazas, así que ¿por qué no reutilizamos el proceso de construcción?

## Objetivos

Un patrón de diseño Builder intenta: 

- Crear un objeto paso a paso rellenando sus campos y creando los objetos incrustados
- Reutilizar el algoritmo de creación de objetos entre muchos objetos

## Implementación

[Ejemplo](./../examples/Builder/creational_test.go)

