import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const ProjectTable = () => {
  return (
    <>
      <Button size="sm">New</Button>
      <Table>
        <TableCaption>项目列表</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">业务线</TableHead>
            <TableHead>项目</TableHead>
            <TableHead>应用</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow>
            <TableCell className="font-medium">共建中台</TableCell>
            <TableCell>库存中台</TableCell>
            <TableCell>miniso-invc-tob-test</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </>
  );
};

export default ProjectTable;
