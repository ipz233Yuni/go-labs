#include <stdbool.h>

double get_price(int material, int glass_type) {
    // 0: Дерево, 1: Метал, 2: Металопластик
    // 0: Однокамерний, 1: Двокамерний
    double prices[3][2] = {
        {2.5, 3.0},   // Дерево
        {0.5, 1.0},   // Метал
        {1.5, 2.0}    // Металопластик
    };
    return prices[material][glass_type];
}

double calc_total(double width, double height, int material, int glass_type, bool has_sill) {
    double area = width * height;
    double price_per_cm2 = get_price(material, glass_type);
    double total = area * price_per_cm2;
    if (has_sill)
        total += 350.0;
    return total;
}
