# Singleton

## 📖 Descripción

El patrón Singleton es fácil de recordar. Como su nombre indica, le proporcionará una única instancia de un objeto y le garantizará que no hay duplicados.
En la primera llamada para usar la instancia, se crea, y luego se reutiliza entre todas las partes de la aplicación que necesitan usar ese comportamiento en particular.

## 💥 Problema

Este patrón resuelve dos problemas al mismo tiempo:

1. Se asegura de que una estructura tenga solo una instancia.

2. Provee un punto de acceso global a la instancia creada.

## ✔️ Solución

Todas las implementaciones del Singleton tienen estos dos pasos en común:

- Hacer que las instancias de la estructura solo se puedan crear a través de un constructor.

- Crear un método para obtener la instancia creada y así todas las siguientes llamadas a este método devuelven el objeto almacenado en caché.

## 🚩 Objetivos

Como guía general, consideramos el uso del patrón Singleton cuando se aplica la siguiente regla:

- Necesitamos un valor único, compartido, de algún tipo particular

- Necesitamos restringir la creación de objetos de algún tipo a una sola unidad a lo largo de todo el programa

## ❔ Cuando aplicarlo

- Cuando quieres usar la misma conexión a una base de datos para hacer querys

- Cuando abrimos por ejemplo una conexión segura (SSH) a un servidor y queremos mandar comandos a través de ella

- Si necesitamos limitar el acceso a alguna variable

- Si necesitamos limitar el número de llamadas en algunos lugares

## 👥 Relación con otros patrones

- Una estructura Facade puede transformarse a menudo en un Singleton, ya que un solo objeto de Facade es suficiente en la mayoría de los casos.

- El patrón Flyweight se asemejaría a Singleton si de alguna manera se lograra reducir todos los estados compartidos de los objetos a un solo objeto de Flyweight. Pero hay dos diferencias fundamentales entre estos patrones:

  1. Debería haber una sola instancia Singleton, mientras que una estructura Flyweight puede tener múltiples instancias con diferentes estados intrínsecos.

  2. El objeto Singleton puede ser mutable. Los objetos de Flyweight son inmutables.

## ⚡️Implementación

[Ejemplo](./examples/Singleton/)
