# Patrones de creación

| Patrón                                         | Descripción                                                                                                                                              |
| ---------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Singleton](./pages/Singleton.md)              | Tener una única instancia de un tipo en todo el programa                                                                                                 |
| [Builder](./pages/Builder.md)                  | Permite separar la construcción de un objeto de su representación con el fin de que el mismo proceso de creación pueda crear diferentes representaciones |
| [Factory](./pages/Factory.md)                  | Delegar la creación de tipos específicos                                                                                                                 |
| [Abstract factory](./pages/AbstractFactory.md) | Proporciona una interface para crear familias de objetos relacionados o dependientes sin especificar sus estructuras concretas.                          |
| [Prototype](./pages/Prototype.md)              | Usa una instancia ya creada de algún tipo para clonarla y completarla con las necesidades particulares de cada contexto                                  |

# Patrones estructurales

| Patrón    | Descripción                                                                                                                                  |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Composite | Proporciona una manera de construir objetos complejos a partir de objetos mas simples                                                        |
| Adapter   | Permite adaptar dos interfaces incompatibles entre sí                                                                                        |
| Bridge    | Permite separar el comportamiento de la representación, permite combinar comportamientos con representaciones                                |
| Proxy     | Tiene como propósito proporcionar un intermediario de un objeto para controlar su acceso                                                     |
| Facade    | Permite ocultar la complejidad de un sistema proveyendo una interfaz sencilla para el cliente                                                |
| Decorator | Proporciona una manera de añadir responsabilidad a un objeto dinámicamente sin afectar el comportamiento original                            |
| Flyweight | También conocido como caché, proporciona un modo de compartir memoria entre partes comunes de objetos para así mejorar el uso de memoria ram |

# Patrones de comportamiento

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

# Patrones de concurrencia

| Patrón                 | Descripción                                                                                                                                                                                                                              |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Barrier                | Su propósito es simple: poner una barrera para que nadie pase hasta que tengamos todos los resultados que necesitamos, algo bastante común en aplicaciones concurrentes.                                                                 |
| Future                 | Es una forma rápida y fácil de lograr estructuras concurrentes para la programación asíncrona. Aprovecharemos las ventajas de las first class functions en Go para desarrollar Futuros.                                                  |
| Pipeline               | Usa la salida de una función como entrada de otra                                                                                                                                                                                        |
| Workers / Pool         | Las goroutines son ligeras, pero el trabajo que realizan puede ser muy pesado. Wokers/pool nos ayuda a resolver este problema, limitando la cantidad de Goroutines disponibles y así controlar los recursos de una manera mas exhaustiva |
| Publisher / Subscriber | Este patrón permite separar generadores de eventos (Publishers) de los que lo usan (Suscribers), con este patrón podríamos crear funcionalidades como hot reload de configuraciones, etc.                                                |

# Fuentes

- [Patrón de diseño](https://es.wikipedia.org/wiki/Patr%C3%B3n_de_dise%C3%B1o)
- [Antipatrón de diseño](https://es.wikipedia.org/wiki/Antipatr%C3%B3n_de_dise%C3%B1o)
