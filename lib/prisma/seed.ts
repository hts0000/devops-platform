import db from "@/lib/prisma/db";
import { Prisma } from "@prisma/client";

const userData: Prisma.UserCreateInput[] = [
  {
    uid: "50009999",
    name: "aaaaaa",
    email: "aaaaaa@miniso.com",
  },
  {
    uid: "50001000",
    name: "bbbbbb",
    email: "bbbbbb@miniso.com",
  },
  {
    uid: "99999999",
    name: "cccccc",
    email: "cccccc@miniso.com",
  },
];

export const userSeed = async () => {
  db.user.createMany({
    data: userData,
  });
};

const seed = async () => {
  await userSeed();
};

seed();
