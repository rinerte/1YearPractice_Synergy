public class Circle : Figure {
    public override string Name => "Окружность";
    public override double GetArea(){
        return 2;
    }
    public override double GetPerimeter(){
        return 2;
    }
}