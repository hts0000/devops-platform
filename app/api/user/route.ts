import { getUsers } from "@/lib/prisma/user/user";
import { NextRequest, NextResponse } from "next/server";

export const GET = async (req: NextRequest) => {
  const size = req.nextUrl.searchParams.get("size");

  console.log(`size: ${size}`);

  const data = await getUsers(0);

  console.log(`data: ${data}`);

  return NextResponse.json(data);
};
