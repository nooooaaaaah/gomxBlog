CREATE MIGRATION m1zu2jsqirclq7fpwuz2baflk2rrozy2ylecvqlns2lu4zp22dpw3q
    ONTO m1msjdb44mst2n5gm6dp2dvml56lh27qgzvgvdsera5xbkgmxv2nkq
{
  ALTER TYPE default::Post {
      CREATE REQUIRED PROPERTY description: std::str {
          SET REQUIRED USING (<std::str>{'descrip'});
      };
      CREATE REQUIRED PROPERTY link: std::str {
          SET REQUIRED USING (<std::str>{'12'});
      };
  };
};
