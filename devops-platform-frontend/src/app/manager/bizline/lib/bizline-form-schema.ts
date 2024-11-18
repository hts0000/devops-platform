import { z } from "zod";

export const BizLineFormSchema = z.object({
  name: z
    .string()
    .regex(
      /^[A-Za-z0-9_-]{3,18}$/,
      "非法的一级项目名称，只能由3~18个字母、数字、下划线、横杠组成"
    ),
  responsibleId: z.string(),
  description: z.string().max(1024, {
    message: "超过1024最大长度",
  }),
});

export type BizLineFormType = z.infer<typeof BizLineFormSchema>;
