CREATE MIGRATION m1vbsizq72jmdamwvi3skqlhvwrbiobp6zkuuei5cfdwhz53evhzta
    ONTO initial
{
  CREATE TYPE default::Category {
      CREATE REQUIRED PROPERTY name: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  CREATE TYPE default::Post {
      CREATE MULTI LINK categories: default::Category;
      CREATE REQUIRED PROPERTY content: std::str;
      CREATE REQUIRED PROPERTY published_on: std::datetime {
          SET default := (std::datetime_current());
      };
      CREATE REQUIRED PROPERTY title: std::str;
  };
  ALTER TYPE default::Category {
      CREATE MULTI LINK posts := (.<categories[IS default::Post]);
  };
};
