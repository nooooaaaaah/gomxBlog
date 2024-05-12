module default {
    type Post {
        overloaded required id: uuid;
        required title: str;
        required content: str;
        required description: str;
        required link: str;
        required published_on: datetime {
            default := datetime_current();
        }
        multi categories: Category;
    }

    type Category {
        overloaded required id: uuid;
        required name: str {
            constraint exclusive;
        }
        multi posts := .<categories[is Post];
    }
};
