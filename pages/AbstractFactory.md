# Abstract Factory

## Descripción

El patrón de diseño Abstract Factory es una nueva capa de agrupamiento para conseguir un objeto compuesto más grande (y más complejo), que se utiliza a través de sus interfaces. La idea detrás de agrupar objetos en familias y agrupar familias es tener grandes fábricas que puedan ser intercambiables y puedan crecer más fácilmente. En las primeras etapas de desarrollo, también es más fácil trabajar con fábricas y fábricas abstractas que esperar hasta que todas las implementaciones concretas estén hechas para comenzar su código. Además, no usará el patrón Abstract Factory desde el principio a menos que sepa que el inventario de su objeto para un campo en particular va a ser muy grande y podría ser fácilmente agrupado en familias.

Este patrón se utiliza comúnmente en muchas aplicaciones y bibliotecas, como las bibliotecas GUI multiplataforma. Piense en un botón, un objeto genérico y una fábrica de botones que le proporciona una fábrica para botones de Microsoft Windows mientras que tiene otra fábrica para botones de Mac OS X. No quieres ocuparte de los detalles de implementación de cada plataforma, sino que sólo quieres implementar las acciones para algún comportamiento específico planteado por un botón.

Además, hemos visto las diferencias al abordar el mismo problema con dos soluciones diferentes: Abstract Factory y el patrón Builder. Como has visto, con el patrón Builder, teníamos una lista no estructurada de objetos (coches con motocicletas en la misma fábrica). Además, alentamos la reutilización del algoritmo de construcción en el patrón Builder. En el patrón Abstract Factory, tenemos una lista muy estructurada de vehículos (la fábrica de motos y la de coches). Tampoco mezclamos la creación de coches con motos, lo que proporciona más flexibilidad en el proceso de creación. Tanto la fábrica de resúmenes como los patrones de construcción pueden resolver el mismo problema, pero sus necesidades particulares le ayudarán a encontrar las pequeñas diferencias que deberían llevarle a tomar una solución u otra.


## Objetivos
Agrupar familias de objetos relacionados es muy conveniente cuando su número de objeto está creciendo tanto que crear un punto único para obtenerlos todos parece la única manera de obtener la flexibilidad de la creación de objetos en tiempo de ejecución. Los siguientes objetivos del método Abstract Factory deben ser claros para usted:
    
    - Proporcionar una nueva capa de encapsulación para los métodos Factory que devuelven una interfaz común para todas las fábricas. 
    
    - Agrupar fábricas comunes en una super Fábrica (también llamada fábrica de fábricas)

## Implementación

[Ejemplo](./../examples/AbstractFactory/)