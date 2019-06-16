# Factory

## Descripción

El patrón del método Factory (o simplemente Factory) es probablemente el segundo patrón de diseño más conocido y utilizado en la industria. Su propósito es abstraer al usuario del conocimiento de la estructura que necesita lograr para un propósito específico, como la recuperación de algún valor, tal vez de un servicio web o una base de datos. El usuario sólo necesita una interfaz que le proporcione este valor. Al delegar esta decisión a una Fábrica, esta Fábrica puede proporcionar una interfaz que se ajuste a las necesidades del usuario. También facilita el proceso de degradación o mejora de la implementación del tipo subyacente si es necesario

Al utilizar el patrón de diseño del método Factory, obtenemos una capa adicional de encapsulación para que nuestro programa pueda crecer en un entorno controlado. 

Con el método Factory, delegamos la creación de familias de objetos a un paquete u objeto diferente para abstraernos del conocimiento del conjunto de posibles objetos que podríamos utilizar. Imagine que quiere organizar sus vacaciones utilizando una agencia de viajes. Usted no se preocupa de hoteles, vuelos, etc. sólo le dice a la agencia el destino que le interesa para que le proporcione todo lo que necesario. La agencia de viajes representa una Fábrica de viajes.

## Objetivos

- Delegar la creación de nuevas instancias de estructuras en una parte diferente del programa. 

- Trabajar a nivel de interfaz en lugar de con implementaciones concretas 

- Agrupación de familias de objetos para obtener un creador de objetos de familia

## Implementación