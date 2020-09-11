# Builder

## 📖 Descripción

El patrón Builder nos ayuda a construir objetos complejos sin instanciar directamente su estructura, ni escribir la lógica que requieren.

## 💥 Problema

Imagina un objeto que podría tener docenas de campos que son estructuras más complejas y anidadas. Tal código de inicialización suele estar encapsulado dentro de un constructor monstruoso con muchos parámetros. O peor aún utilizando estructuras para cada caso, por ejemplo imaginemos que queremos construir un coche o una moto, una solución podría ser crear estructuras para cada uno con sus constructores, pero es que la mayoría de propiedades serían las mismas.

## ✔️ Solución

La solucióm a este problema es utilizar el patrón Builder que nos aconseja extraer la construcción del objeto en una estructura separada que será la encargada de construir el objeto (Builder) y una estructura que será la encargada de dirigir la construcción (Director). El patrón organiza la creación del objeto en diferentes pasos por lo que no necesitamos tener un mega constructor. La parte mas importante es que no necesitas llamar a todos los pasos para generar el objeto.

## 🚩 Objetivos

Un patrón de diseño Builder intenta:

- Crear un objeto paso a paso rellenando sus campos y creando los objetos incrustados

- Reutilizar el algoritmo de creación de objetos entre muchos objetos

## ❔ Cuando aplicarlo

- Usar el patrón Builder para deshacerse de un "constructor monstruoso".

- Utilice el patrón Builder cuando quiera que su código sea capaz de crear diferentes representaciones de algún producto (por ejemplo, casas de piedra y madera).

- Usar el patrón Builder para construir árboles compuestos u otros objetos complejos.

## 👥 Relación con otros patrones

- Abstract Factory, Builder y Prototype pueden ser implementados como Singletons

- Puedes usar Builder al crear árboles compuestos complejos porque puedes programar sus pasos de construcción para que funcionen de forma recursiva.

- El patrón Builder se centra en construir objetos complejos paso a paso. Abstract Factory se especializa en crear familias de objetos relacionados. Abstract Factory devuelve el objeto inmediatamente, mientras que Builder le permite ejecutar algunos pasos de construcción adicionales antes de generar el objeto.

- Puedes combinar el patrón Builder con el patrón Bridge: la estructura de director juega el papel de la abstracción, mientras que los diferentes constructores actúan como implementaciones.

## ⚡️Implementación

[Ejemplo](./examples/Builder/)
