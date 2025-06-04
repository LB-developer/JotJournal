"use client";
import { useEstablishUser } from "@/hooks/user";
import { updateJotCompletion } from "@/lib/jots/updateJotCompletion";
import { Jot, JotCollection } from "@/types/jotTypes";
import { useState } from "react";
import { Checkbox } from "../ui/checkbox";

interface Props {
    jotCollection: JotCollection;
}

export default function JotDisplay({ jotCollection }: Props) {
    const [jots, setJots] = useState<JotCollection>(jotCollection);
    useEstablishUser();

    const handleUpdateJot = async (
        e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
        jotToUpdate: Jot,
    ): Promise<void> => {
        e.preventDefault();

        const jotID = jotToUpdate.id;
        const habit = jotToUpdate.habit;
        const changedCompletion = !jotToUpdate.isCompleted;

        // TODO: create tag for caching
        //   structure: ["id-habit1", "id-habit2", etc...] for as many habits that are in jots
        //
        // const tags: string[] = []
        // for (const key of Object.keys(jots)) {
        //   tags.push(`tag-${context.user.ID}-${key}`);
        // }

        // update jot in db
        await fetch(`/api/jots/`, {
            method: "PATCH",
            body: JSON.stringify({ jotID, isCompleted: changedCompletion }),
        });

        // update jot locally
        const updatedJots = { ...jots };
        const length = updatedJots[habit].length - 1;
        updateJotCompletion(updatedJots[habit], 0, length, jotID);
        setJots(updatedJots);
    };

    return (
        <section className="flex flex-row gap-2 p-4 bg-gray-400">
            {Object.entries(jots).map(([habit, jots]) => (
                <div key={habit}>
                    <h4 className="text-sm font-semibold mb-1 pb-2 -rotate-45">
                        {habit}
                    </h4>
                    <div className="flex flex-col">
                        {jots.map((jot) => (
                            <Checkbox
                                key={jot.id}
                                className={`w-5 h-5 rounded-sm border ${
                                    jot.isCompleted
                                        ? "bg-green-500 border-green-700"
                                        : "bg-gray-200 border-gray-400"
                                }`}
                                title={new Date(jot.date).toDateString()}
                                onClick={(e) => handleUpdateJot(e, jot)}
                            />
                        ))}
                    </div>
                </div>
            ))}
        </section>
    );
}
