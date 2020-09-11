# Factory

## 📖 Descripción

El propósito del patrón Factory Method (o simplemente Factory) es abstraer al usuario del conocimiento de la estructura que necesita lograr para un propósito específico, como la recuperación de algún valor, tal vez de un servicio web o una base de datos. El usuario sólo necesita una interfaz que le proporcione este valor. Al delegar esta decisión a una Fábrica, esta Fábrica puede proporcionar una interfaz que se ajuste a las necesidades del usuario. También facilita el proceso de degradación o mejora de la implementación del tipo subyacente si es necesario

Al utilizar el patrón de diseño del método Factory, obtenemos una capa adicional de encapsulación para que nuestro programa pueda crecer en un entorno controlado.

Con el método Factory, delegamos la creación de objetos a un paquete u objeto diferente para abstraernos del conocimiento del conjunto de posibles objetos que podríamos utilizar. Imagine que quiere organizar sus vacaciones utilizando una agencia de viajes. Usted no se preocupa de hoteles, vuelos, etc. sólo le dice a la agencia el destino que le interesa para que le proporcione todo lo que necesario. La agencia de viajes representa una Fábrica de viajes.

## 💥 Problema

Imagina queremos crear una aplicación de ecommerce. En la primera versión de la App tenemos un único método de pago que es en cash, por lo que los programadores creamos una entidad llamada CashPayment la cual contiene el método pay que realiza toda la funcionalidad a la hora de pagar. Como el negocio va bien, nuestros stakeholders nos comentan que en el mercado esta muy demandado el pago con tarjeta y que probablemente en un futuro queramos poder pagar con paypal, como el código está altamente acoplado a la estructura CashPayment para añadir el pago con tarjeta deberíamos duplicar toda la lógica común en ambos sitios.

## ✔️ Solución

El patrón Factory Method dice que reemplacemos la construcción del objeto directamente con llamadas a un factory method. En el ejemplo anterior creariamos una función GetPaymentMethod que se encargará de elegir el método de pago que se va a utilizar a la hora de pagar y en las funciones o estructuras con la lógica de negocio utilizarán la interface del PaymentMethod que sería la abstracción para desacoplarnos del método de pago.

## 🚩 Objetivos

- Delegar la creación de nuevas instancias de estructuras en una parte diferente del programa.

- Trabajar a nivel de interfaz en lugar de con implementaciones concretas

- Agrupación de familias de objetos para obtener un creador de objetos de familia

## ❔ Cuando aplicarlo

- Se puede utilizar cuando no se tiene claro de antemano los tipos de dependencias que tendrán los objetos.
- Cuando quieras proporcionar una manera de extender el dominio
- Cuando quieras ahorrar recursos del sistema reutilizando objetos existentes en vez de reconstruirlos cada vez

## 👥 Relación con otros patrones

- Abstract factory
- Builder
- Prototype
- Template method

## ⚡️Implementación

[Ejemplo](./examples/Factory/)
