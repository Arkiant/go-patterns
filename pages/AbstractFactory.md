# Abstract Factory

## 📖 Descripción

El patrón de diseño Abstract Factory es una nueva capa de agrupamiento para conseguir un objeto compuesto más grande (y más complejo), que se utiliza a través de sus interfaces. La idea detrás de agrupar objetos en familias y agrupar familias es tener grandes fábricas que puedan ser intercambiables y puedan crecer más fácilmente. En las primeras etapas de desarrollo, también es más fácil trabajar con fábricas y fábricas abstractas que esperar hasta que todas las implementaciones concretas estén hechas para comenzar su código. 

Este patrón se utiliza comúnmente en muchas aplicaciones y bibliotecas, como las bibliotecas GUI multiplataforma. Piense en un botón, un objeto genérico y una fábrica de botones que le proporciona una fábrica para botones de Microsoft Windows mientras que tiene otra fábrica para botones de Mac OS X. No quieres ocuparte de los detalles de implementación de cada plataforma, sino que sólo quieres implementar las acciones para algún comportamiento específico planteado por un botón.

## 💥 Problema

Imaginemos que estamos creando una aplicación de renting, en ella tenemos Motos y Coches, de estos tenemos diferentes modelos, por ejemplo de moto tenemos: deportiva y de camino, y de coches tenemos de lujo y familiar. Pero si tenemos el código muy acoplado podremos tener problemas a la hora de organizar el código y hacerlo más mantenible.

## ✔️ Solución

Lo primero que sugiere el patrón Abstract Factory es declarar explícitamente las interfaces para cada producto distinto de la familia de productos (por ejemplo, Motorbike y Car). Luego se puede hacer que todas las variantes de productos sigan esas interfaces. Por ejemplo, todas las variantes de moto y coche podrían implementar esa interfaz.

El siguiente paso es declarar la Abstract Factory, una interfaz con una lista de métodos de creación para todos los productos que forman parte de la familia de productos, en el ejemplo el método Build hace esta función. Estos métodos deben devolver los tipos de productos abstractos representados por las interfaces que hemos extraído anteriormente: Car y Motorbike sucesivamente.

Ahora, ¿qué hay de las variantes del producto? Para cada variante de una familia de productos, creamos una estructura Fabric separada basada en la interfaz VehicleFactory. Una fábrica es una estructura que devuelve productos de un tipo particular. Por ejemplo, la Fábrica de Motos sólo puede crear objetos de Motos de carretera o deportivas.

El código del cliente tiene que trabajar con ambas fábricas y productos a través de sus respectivas interfaces abstractas. Esto permite cambiar el tipo de fábrica que se pasa al código de cliente, así como la variante de producto que recibe el código de cliente, sin romper el código de cliente real.

## 🚩 Objetivos

Agrupar familias de objetos relacionados es muy conveniente cuando su número de objeto está creciendo tanto que crear un punto único para obtenerlos todos parece la única manera de obtener la flexibilidad de la creación de objetos en tiempo de ejecución. Los siguientes objetivos del método Abstract Factory deben ser claros para usted:
    
- Proporcionar una nueva capa de encapsulación para los métodos Factory que devuelven una interfaz común para todas las fábricas. 

- Agrupar fábricas comunes en una super Fábrica (también llamada fábrica de fábricas)

## ❔ Cuando aplicarlo

-  Utilice el patrón Abstract Factory cuando su código necesite trabajar con varias familias de productos relacionados, pero no quiere que dependa de las estrcuturas concretas de esos productos, ya que podrían ser desconocidas de antemano o simplemente quiere permitir una futura extensibilidad.

## 👥 Relación con otros patrones

- El patrón Abstract Factory puede servir como alternativa al patrón Facade cuando sólo se quiere ocultar la forma en que se crean los objetos del subsistema a partir del código del cliente.

## ⚡️Implementación

[Ejemplo](./../examples/AbstractFactory/)