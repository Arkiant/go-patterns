# Prototype

---

## 📖 Descripción

El objetivo del patrón Prototype es tener un objeto o conjunto de objetos que ya esté creado en el momento de la compilación, pero que se puede clonar tantas veces como se desee en tiempo de ejecución. Esto es útil, por ejemplo, como plantilla predeterminada para un usuario que acaba de registrarse en su página web o como plan de precios predeterminado en algún servicio. La diferencia clave entre esto y un patrón Builder es que los objetos se clonan para el usuario en lugar de construirlos en tiempo de ejecución. También puede crear una solución similar a una caché, almacenando información utilizando un prototipo.

## 💥 Problema

## ✔️ Solución

## 🚩 Objetivos

El objetivo principal del patrón de diseño Prototype es evitar la creación repetitiva de objetos. Imagine un objeto por defecto compuesto de docenas de campos y tipos incrustados. No queremos escribir todo lo que necesita este tipo cada vez que usamos el objeto, especialmente si podemos estropearlo creando instancias con diferentes fundamentos:

- Mantener un conjunto de objetos que serán clonados para crear nuevas instancias

- Proporcionar un valor por defecto de algún tipo para empezar a trabajar sobre él.

- CPU libre de inicialización de objetos complejos para tomar más recursos de memoria

## ❔ Cuando aplicarlo

## 👥 Relación con otros patrones

## ⚡️Implementación

[Ejemplo](./examples/Prototype/)
