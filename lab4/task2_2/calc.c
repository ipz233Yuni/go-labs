#include <stdbool.h>

double get_tour_price(int country, int season) {
    // 0: Болгарія, 1: Німеччина, 2: Польща
    // 0: Літо, 1: Зима
    double prices[3][2] = {
        {100, 150},
        {160, 200},
        {120, 180}
    };
    return prices[country][season];
}

double calc_total(double days, int country, int season, bool guide, bool luxury) {
    double base = get_tour_price(country, season);
    double total = base * days;
    if (guide) total += 50 * days;
    if (luxury) total *= 1.2;
    return total;
}
