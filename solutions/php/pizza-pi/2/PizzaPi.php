<?php

class PizzaPi
{
    public function calculateDoughRequirement($pizzas, $serves)
    {
        return $pizzas * (($serves * 20) + 200);
    }

    public function calculateSauceRequirement($pizzas, $sauce)
    {
        return $pizzas / ($sauce / 125);
    }

    public function calculateCheeseCubeCoverage($cheese, $thickness, $pizza_diameter)
    {
        return floor(($cheese**3) / ($thickness * M_PI * $pizza_diameter));
    }

    public function calculateLeftOverSlices($pizza, $friends)
    {
        return ($pizza * 8) % $friends;
    }
}
