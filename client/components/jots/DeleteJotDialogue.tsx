import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { Jot } from "@/types/jotTypes";
import { XIcon } from "lucide-react";

interface Props {
    jot: Jot;
    submit: (
        e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
        jot: Jot,
        action: "delete",
    ) => Promise<void>;
}

export function DeleteJotDialogue({ jot, submit }: Props) {
    return (
        <Dialog>
            <div className="flex flex-row align-center max-h-4">
                <DialogTrigger asChild>
                    <Button
                        variant={"ghost"}
                        className="whitespace-nowrap self-center"
                    >
                        <XIcon />
                    </Button>
                </DialogTrigger>
                <h3 className="self-center">{jot.habit}</h3>
            </div>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Delete {jot.habit}</DialogTitle>
                    <DialogDescription>
                        Are you sure you want to delete this Jot? This action
                        cannot be undone.
                    </DialogDescription>
                </DialogHeader>
                <DialogFooter>
                    <DialogClose asChild>
                        <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button
                        variant="outline"
                        onClick={(e) => submit(e, jot, "delete")}
                    >
                        Delete
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
