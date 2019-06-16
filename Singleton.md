# Singleton

## Descripción

El patrón Singleton es fácil de recordar. Como su nombre indica, le proporcionará una única instancia de un objeto y le garantizará que no hay duplicados. 
En la primera llamada para usar la instancia, se crea, y luego se reutiliza entre todas las partes de la aplicación que necesitan usar ese comportamiento en particular.

Utilizaremos el patrón Singleton en diferentes situaciones, por ejemplo:

- Cuando quieres usar la misma conexión a una base de datos para hacer querys
- Cuando abrimos por ejemplo una conexión segura (SSH) a un servidor y queremos mandar comandos a través de ella
- Si necesitamos limitar el acceso a alguna variable
- Si necesitamos limitar el número de llamadas en algunos lugares

Las posibilidades son infinitas, solo mencionamos algunas de ellas.

## Objetivos

Como guía general, consideramos el uso del patrón Singleton cuando se aplica la siguiente regla:

- Necesitamos un valor único, compartido, de algún tipo particular
- Necesitamos restringir la creación de objetos de algún tipo a una sola unidad a lo largo de todo el programa