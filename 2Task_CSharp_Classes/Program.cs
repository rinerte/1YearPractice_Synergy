// See https://aka.ms/new-console-template for more information
Console.WriteLine("Hello, World!");

List<Figure> figures = new List<Figure>{
    new Triangle(5),
    new Triangle(3,4,5)
};

foreach (var figure in figures){
    Console.WriteLine($"Фигура: {figure.Name}");
    Console.WriteLine($"Площадь: {figure.GetArea():F2}");
    var perimimeter = figure is Circle ? "Длинна окружности" : "Периметр";
    Console.WriteLine($"{perimimeter}: {figure.GetPerimeter():F2} \n");
}

