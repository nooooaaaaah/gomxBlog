CREATE MIGRATION m1msjdb44mst2n5gm6dp2dvml56lh27qgzvgvdsera5xbkgmxv2nkq
    ONTO m14b5hx2fo4a3gl7bnowlxclt72zvm5e56tkwhpcptsttfzwi7bsva
{
  ALTER TYPE default::Post {
      ALTER PROPERTY id {
          SET OWNED;
          SET REQUIRED;
          SET TYPE std::uuid;
      };
  };
  ALTER TYPE default::Category {
      ALTER PROPERTY id {
          SET OWNED;
          SET REQUIRED;
          SET TYPE std::uuid;
      };
  };
};
