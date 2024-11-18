import Long from "long";
import { z } from "zod";

export const BizLineEntitySchema = z.object({
  // 将Long类型转换成string
  // 第二个参数最终验证转换是否成功
  id: z.preprocess(
    (val) => (Long.isLong(val) ? val.toString() : val),
    z.string()
  ),
  bizLine: z.object({
    name: z.string(),
    responsibleId: z.preprocess(
      (val) => (Long.isLong(val) ? val.toString() : val),
      z.string()
    ),
    description: z.string(),
  }),
});

export type BizLineEntity = z.infer<typeof BizLineEntitySchema>;
