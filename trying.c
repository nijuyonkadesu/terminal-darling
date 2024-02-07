#include <notcurses/notcurses.h>

int main() {
    struct notcurses_options opts = {
        .flags = NCOPTION_INHIBIT_SETLOCALE | NCOPTION_NO_ALTERNATE_SCREEN,
    };

    struct notcurses* nc = notcurses_init(&opts, NULL);
    if (!nc) {
        return -1;
    }

    struct ncplane* stdplane = notcurses_stdplane(nc);
    
    ncplane_move_yx(stdplane, 100, 100);

    if(ncplane_putstr(stdplane, "24k") < 0 ) {
        notcurses_stop(nc);
        return -1;
    }

    if (notcurses_render(nc)) {
        notcurses_stop(nc);
        return -1;
    }

    notcurses_get(nc, NULL, NULL);

    notcurses_stop(nc);
    return 0;
}
