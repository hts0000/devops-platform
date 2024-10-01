import { getUser } from "@/lib/prisma/user/user";
import { NextRequest, NextResponse } from "next/server";

export const GET = async (
  req: NextRequest,
  { params }: { params: { uid: string } }
) => {
  const data = await getUser(params.uid);

  return NextResponse.json(data);
};
