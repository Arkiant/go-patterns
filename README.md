# :wave: Introducción

Este repositorio me sirve para guardar todo el conocimiento sobre patrones, arquitecturas, antipatrones, refactoring y cosas avanzadas. Decidí crear el repositorio opensource para poder reunir en un solo lugar todo el conocimiento que voy adquiriendo en mi carrera profesional sobre Go, arquitectura, sistemas distribuídos, y compartirlo con todo el mundo en español.

Cualquiera es bienvenido a aportar sus conocmientos y de esa manera podamos aprovechar y aprender todos, incluído yo.

Iré actualizando a menudo el repositorio con todas las notas que tengo guardadas en papel.

# :hammer: Patrones de diseño

## Patrones de creación

| Patrón                                                  | Descripción                                                                                                                                              |
| ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Singleton](./pages/patterns/Singleton.md)              | Tener una única instancia de un tipo en todo el programa                                                                                                 |
| [Builder](./pages/patterns/Builder.md)                  | Permite separar la construcción de un objeto de su representación con el fin de que el mismo proceso de creación pueda crear diferentes representaciones |
| [Factory](./pages/patterns/Factory.md)                  | Delegar la creación de tipos específicos                                                                                                                 |
| [Abstract factory](./pages/patterns/AbstractFactory.md) | Proporciona una interface para crear familias de objetos relacionados o dependientes sin especificar sus estructuras concretas.                          |
| [Prototype](./pages/patterns/Prototype.md)              | Usa una instancia ya creada de algún tipo para clonarla y completarla con las necesidades particulares de cada contexto                                  |

## Patrones estructurales

| Patrón    | Descripción                                                                                                                                  |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Composite | Proporciona una manera de construir objetos complejos a partir de objetos mas simples, en Go, es lo más parecido a herencia que tenemos      |
| Adapter   | Permite adaptar dos interfaces incompatibles entre sí                                                                                        |
| Bridge    | Permite separar el comportamiento de la representación, permite combinar comportamientos con representaciones                                |
| Proxy     | Tiene como propósito proporcionar un intermediario de un objeto para controlar su acceso                                                     |
| Facade    | Permite ocultar la complejidad de un sistema proveyendo una interfaz sencilla para el cliente                                                |
| Decorator | Proporciona una manera de añadir responsabilidad a un objeto dinámicamente sin afectar el comportamiento original                            |
| Flyweight | También conocido como caché, proporciona un modo de compartir memoria entre partes comunes de objetos para así mejorar el uso de memoria ram |

## Patrones de comportamiento

| Patrón                  | Descripción                                                                                                                                                                                                                                                     |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Chain of Responsability | Permite pasar la responsabilidad entre diferentes handlers, al recibir la solicitud, cada handler decide si procesa la solicitud o la pasa al siguiente handler de la cadena                                                                                    |
| Command                 | Este patrón permite solicitar una operación a un objeto sin conocer realmente el contenido de esta operación, ni el receptor real de la misma. Para ello se encapsula la petición como un objeto, con lo que además facilita la parametrización de los métodos. |
| Template Method         | Permite inyectar en una implementación una o varias funciones particulares en el esqueleto de un algoritmo.                                                                                                                                                     |
| Memento                 | Nos permite guardar y restaurar el estado previo de un objeto sin revelar los detalles de la implementación                                                                                                                                                     |
| Visitor                 | Te permite separar los algoritmos de los objetos en los que operan.                                                                                                                                                                                             |
| State                   | Se utiliza cuando el comportamiento de un objeto cambia dependiendo del estado del mismo.                                                                                                                                                                       |
| Mediator                | Nos proporciona una manera de reducir dependencias caóticas entre objetos. Este patrón restringe la comunicación directa entre objetos y los fuerza a colaborar solo vía el mediador                                                                            |
| Iterator                | Permite recorrer los elementos de una colección sin exponer su representación subyacente.                                                                                                                                                                       |
| Observer                | Nos permite notificar a varios objetos de cualquier evento que sucede y al que dichos objetos están subscritos.                                                                                                                                                 |
| Strategy                | Permite definir una familia de algoritmos, poner cada uno de ellos en una clase separada, y hacer que sus objetos sean intercambiables.                                                                                                                         |

## Patrones de concurrencia

| Patrón                 | Descripción                                                                                                                                                                                                                              |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Barrier                | Su propósito es simple: poner una barrera para que nadie pase hasta que tengamos todos los resultados que necesitamos, algo bastante común en aplicaciones concurrentes.                                                                 |
| Future                 | Es una forma rápida y fácil de lograr estructuras concurrentes para la programación asíncrona. Aprovecharemos las ventajas de las first class functions en Go para desarrollar Futuros.                                                  |
| Pipeline               | Usa la salida de una función como entrada de otra                                                                                                                                                                                        |
| Workers / Pool         | Las goroutines son ligeras, pero el trabajo que realizan puede ser muy pesado. Wokers/pool nos ayuda a resolver este problema, limitando la cantidad de Goroutines disponibles y así controlar los recursos de una manera mas exhaustiva |
| Publisher / Subscriber | Este patrón permite separar generadores de eventos (Publishers) de los que lo usan (Suscribers), con este patrón podríamos crear funcionalidades como hot reload de configuraciones, etc.                                                |

# :satellite: Patrones de sistemas distribuidos

## Patrones de gestión de multiples contenedores en un solo nodo

| Patrón     | Descripción                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Sidecar    | Este patrón extiende un contenedor con otro contenedor, por ejemplo el contenedor principal puede ser un servidor web y podría ser emparejado con un contenedor que guarde logs.                                                                                                                                                                                                                                            |
| Ambassador | Este patrón hace de proxy hacia el contenedor principal, por ejemplo, un desarrollador podría emparejar una aplicación que este enviando datos por el protocolo de memcache, el contenedor abassador lo que haría es escuchar y podría estar enviando estos datos a un cluster por ejemplo de shard de memcache sin que el contenedor principal lo sepa, dicho contenedor solo sabría que esta enviando datos en localhost. |
| Adapter    | Este patrón es parecido al Abassador pero con la diferencia de que la aplicación expone una interfaz para poder conectar ambos contenedores mediante interfaces. Por ejemplo los adaptadores que aseguran que todos los contenedores en un sistema tienen la misma interfaz de monitorización.                                                                                                                              |

## Patrones de múltiples nodos

| Patrón          | Descripción                                                                                                                                                                                    |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Leader election | Este patrón consiste en tener un contenedor como líder dentro de un conjunto de replicas, de ese modo si el lider se cae, mediante este patrón elegiríamos al nuevo líder.                     |
| Work queue      | Este patrón consiste en encadenar multiples contenedores uniendo la salida de uno con la entrada de otro para encadenarlos y poder reutilizar sus funcionalidades cuando sea necesario         |
| Scatter/Gather  | En ese sistema, un cliente externo envía una solicitud inicial a un nodo "raíz" o "padre". Esta raíz envía la solicitud a un gran número de servidores para realizar los cálculos en paralelo. |

# :roller_coaster: Arquitecturas y paradigmas

| Nombre                                          | Descripción                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Driven Domain Development (DDD)                 | Arquitectura basada en una profunda conexión entre la implementación y los conceptos de negocio, estos conceptos son el dominio. De esta manera podemos desacoplar toda la lógica de negocio                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| Event Driven Architecture (EDA)                 | Es un paradigma de arquitectura de software que se basa en producir, detectar, consumir y reaccionar a eventos                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Command Query Responsability Segregation (CQRS) | Arquitectura basada en separar lecturas (Query) de escrituras (Command), esta arquitectura es muy utilizada en sistemas dónde hay muchas mas búsquedas que commits y de esta manera poder escalar por separado ambas funcionalidades.                                                                                                                                                                                                                                                                                                                                                                                                |
| Microservice Architecture                       | Esta arquitectura de software se basa en separar cada unidad lógica de negocio en un microservicio. Una de las principales ventajas es que se pueden escalar por separado aunque su mayor desventaja es que añade complejidad accidental y una difícil monitorización                                                                                                                                                                                                                                                                                                                                                                |
| Service Oriented Architecture                   | La orientación a servicios es una forma de pensar en servicios, su construcción y sus resultados. Un servicio es una representación lógica de una actividad de negocio que tiene un resultado de negocio específico. Los microservicios son una interpretación moderna de la arquitectura orientada a servicios usada para construir sistemas distribuidos. Los servicios en una arquitectura de microservicios son procesos que se comunican con otros a través de una red para conseguir el objetivo final. Estos servicios pueden usar protocolos simples (típicamente HTTP con REST o mensajería liviana como RabbitMQ o ZeroMQ) |
| Serverless Architecture                         | Este tipo de arquitectura usa backend de terceros como servicio (BaaS), utilizan funciones como servicio (FaaS). Esta arquitectura elimina la necesidad de crear un backend.                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

# :mount_fuji: Antipatrones

:construction: - WORK IN PROGRESS

# :spaghetti: Refactoring

:construction: - WORK IN PROGRESS
