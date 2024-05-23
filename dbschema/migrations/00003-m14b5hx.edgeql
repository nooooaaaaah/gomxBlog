CREATE MIGRATION m14b5hx2fo4a3gl7bnowlxclt72zvm5e56tkwhpcptsttfzwi7bsva
    ONTO m1a6xzpryayn2ocbdheq4hwblxxkxawphnuepviyin4hcocrqox2sq
{
  ALTER TYPE default::Category {
      CREATE MULTI LINK posts := (.<categories[IS default::Post]);
  };
};
