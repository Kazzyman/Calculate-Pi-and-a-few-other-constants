package main

import (
	"math/big"
	"time"
)

// GridPi approximates the value of Pi using a grid-based method with big.Float
/*
	The realization that the area of a circle can be determined from its circumference and diameter, and consequently the value of
	π, has roots in ancient geometry and the study of ratios.

	Here’s a brief historical and mathematical exploration of how this understanding developed:

	Historical Perspective:

	Ancient Greek Geometry:
		The ancient Greeks, particularly mathematicians like Euclid and Archimedes, were deeply interested in the properties of circles.
		They studied the relationships between different aspects of a circle (such as radius, diameter, circumference, and area)
		using the proven geometric methods of that time period.

	Euclid:
		In his work "Elements" Euclid did not explicitly use the value of π, but he did described the properties of circles in terms of ratios
		and geometric constructions.

	Archimedes:
		Archimedes made significant contributions by approximating π and demonstrating that the area of a circle is proportional to
		the square of its radius. To accomplish this he used the method of exhaustion, a precursor to integral calculus, to show that the area
		of a circle is given as: A=Pi*r*r

		Archimedes' Method:

		Archimedes inscribed and circumscribed polygons around a circle to approximate the value of π. By increasing the number of sides,
		he showed that the ratio of the circumference to the diameter approaches a constant value, which we now know as π. He also demonstrated
		that the area of a circle can be thought of as a series of triangles. By comparing the circle to these inscribed and circumscribed polygons,
		he derived that the area A is proportional to r*r , specifically, that A=Pi*r*r

	Mathematical Insight:

		Relationship Between Circumference and Area:

			To find the area 𝐴, consider the fact that a circle can be thought of as being composed of an infinite number of infinitesimally-thin
			concentric rings. The area of each ring is approximately the circumference of the ring times its infinitesimal thickness. When you
			sum these up, you get the total area of the circle -- demonstrating that circumference is proportional to area.

			Specifically, the circumference C of a circle is given by 𝐶=2𝜋𝑟 (or C=2πr), where 𝑟 (r) is the radius.

	Visualizing the Area:

		One way to visualize the area is to imagine "unrolling" the circumference into a series of isosceles triangles. These triangles collectively
		approximate the area of the circle.

		If you cut a circle into many small wedges (isosceles triangles, each with one slightly curved base side) and rearrange them,
		they form a shape that approximates a parallelogram (or a rectangle in the limit). The height of this shape is the radius 𝑟,
		and the base is half the circumference; because we used half of the wedges to complete the solid-formed rectangle.

	Conclusion:

		The understanding that the area of a circle is related to its circumference and diameter arises from the intrinsic properties of circles
		and the geometric explorations of ancient mathematicians. The discovery of π as the ratio of circumference to diameter was crucial because
		it provided a constant that could be used to relate different measurements of a circle. Once π was understood, mathematicians could use it
		to derive other formulas, including the formula for the area of a circle.

		In essence, knowing the precise relationship between the circumference and diameter (through π) allowed mathematicians to establish a
		consistent and accurate method to calculate the area of a circle, illustrating the deep interconnectedness of geometric properties.

	The algorithm set forth below directly converges on Pi by sampling points in a square which circumscribes a quarter circle and then
	determining the proportion of points that fall inside the quarter circle. We will randomly generate points within the square and count
	how many fall inside the quarter circle, and how many fall outside of it; and, then take the ratio of the two. Thus giving the area
	of one quarter of the complete circle. To make it easy, will use a unit square (a square measuring one unit by one unit) thereby
	inscribing a quarter of a circle with a radius of one unit. This procedure will give us an approximation of Pi because each point
	could be the top left corner of an asymptotically (very small) square. What we accomplish is that we fill each region with small
	squares of uniform size to find the area of each region, one of which is the area inside the inscribed circle.

	If you already know the approximate area of a circle with known radius (and we will know it since we used a unit circle, that is
	a circle with a radius of 1) you can then approximate 𝜋 using the equation C=4a/d to find the Circumference. Knowing both the
	Circumference and the Radius makes it trivially easy to then calculate Pi : the ratio of Circumference to diameter.

	Explanation of the classic Monte Carlo method (psudo psudo code):
		1. Initialization: The random number generator is seeded with the current time to ensure different results on each run.
		2. Random Sampling: The rand.Float64() function generates random floating-point numbers between 0 and 1.
		3. Condition Check: The condition x*x + y*y <= 1 checks whether the point (𝑥, 𝑦) lies inside the quarter circle.
		4. Counting and Ratio Calculation: The ratio of points inside the quarter circle to the total points is multiplied
		by 4 to approximate π.

	However, in our algorithm we have elected to forgo using a random number generator. Instead, we will begin by dividing our unit square
	into a grid of smaller squares and then separating them into two sets. One set will be comprised of all those that fall outside the
	inscribed quarter circle, and the other set will be of those squares that fall inside the inscribed quarter circle.
*/
func GridPi(gridSize int) *big.Float {
	start := time.Now()
	insideCircle := big.NewInt(0)
	totalPoints := big.NewInt(int64(gridSize * gridSize))
	increment := big.NewFloat(1.0 / float64(gridSize))

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			x := new(big.Float).Mul(increment, big.NewFloat(float64(i)))
			y := new(big.Float).Mul(increment, big.NewFloat(float64(j)))

			// x*x + y*y
			xSquared := new(big.Float).Mul(x, x)
			ySquared := new(big.Float).Mul(y, y)
			sum := new(big.Float).Add(xSquared, ySquared)

			// Check if x^2 + y^2 <= 1
			if sum.Cmp(big.NewFloat(1.0)) <= 0 {
				insideCircle.Add(insideCircle, big.NewInt(1))
			}
		}
		iterationsForMonte16 = i
	}

	// Calculate Pi approximation
	four := big.NewFloat(4.0)
	insideCircleF := new(big.Float).SetInt(insideCircle)
	totalPointsF := new(big.Float).SetInt(totalPoints)
	piApprox := new(big.Float).Quo(insideCircleF, totalPointsF)
	piApprox.Mul(piApprox, four)

	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	// iters := 99
	printResultStatsLong(piApprox, 0, "MonteCarlo", iterationsForMonte16, TotalRun, selection)

	return piApprox
}
