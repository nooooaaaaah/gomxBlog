CREATE MIGRATION m1a6xzpryayn2ocbdheq4hwblxxkxawphnuepviyin4hcocrqox2sq
    ONTO m1vbsizq72jmdamwvi3skqlhvwrbiobp6zkuuei5cfdwhz53evhzta
{
  ALTER TYPE default::Category {
      DROP LINK posts;
  };
};
