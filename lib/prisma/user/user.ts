import db from "@/lib/prisma/db";

export const getUser = async (uid: string) => {
  return await db.user.findFirstOrThrow({
    select: {
      id: true,
      name: true,
      email: true,
      businessLines: true,
      projects: true,
    },
    where: {
      uid: {
        equals: uid,
      },
    },
  });
};

export const getUsers = async (size: number) => {
  return await db.user.findMany({});
};
