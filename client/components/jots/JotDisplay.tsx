"use client";
import { useEstablishUser } from "@/hooks/user";
import { updateJotCompletion } from "@/lib/jots/updateJotCompletion";
import { Jot, JotCollection } from "@/types/jotTypes";
import { FormEvent, Fragment, useState } from "react";
import { Checkbox } from "../ui/checkbox";
import { CreateJotDialogue } from "./CreateJotDialogue";
import getDaysInMonth from "@/lib/date/daysInMonth";

interface Props {
    jotCollection: JotCollection;
    month: string;
    year: string;
}

export default function JotDisplay({ jotCollection, month, year }: Props) {
    const [jots, setJots] = useState<JotCollection>(jotCollection);
    const daysInMonth = getDaysInMonth(Number(year), Number(month));
    useEstablishUser();

    const today = new Date();

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

    const handleSubmit = async (
        e: FormEvent<HTMLFormElement>,
    ): Promise<void> => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const name = formData.get("name");
        const monthAsNum = Number(month);
        const yearAsNum = Number(year);

        const res = await fetch("/api/jots", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, month: monthAsNum, year: yearAsNum }),
        });

        const newJot: Jot[] = await res.json();

        console.log(newJot);

        const habit = newJot[0].habit;

        setJots({
            [habit]: newJot,
            ...jots,
        });
    };

    return (
        <section
            className="grid w-full gap-1 p-4 text-sm"
            style={{
                gridTemplateColumns: `auto repeat(${daysInMonth}, 1fr)`,
            }}
        >
            {/* Row 1: Day letters */}
            <>
                <CreateJotDialogue submit={handleSubmit} />
                {Array.from({ length: daysInMonth }).map((_, i) => {
                    const date = new Date(
                        today.getFullYear(),
                        today.getMonth(),
                        i + 1,
                    );
                    const letter = date.toLocaleDateString("en-NZ", {
                        weekday: "short",
                    })[0];
                    return (
                        <div
                            key={`day-letter-${i}`}
                            className="text-center font-semibold"
                        >
                            {letter}
                        </div>
                    );
                })}
            </>

            {/* Row 2: Day numbers */}
            <>
                <div /> {/* Second-left empty cell */}
                {Array.from({ length: daysInMonth }).map((_, i) => (
                    <div
                        key={`day-num-${i}`}
                        className="text-center text-muted-foreground"
                    >
                        {i + 1}
                    </div>
                ))}
            </>
            {/* Rows 3+: Habit Rows */}
            {Object.entries(jots).map(([habit, jots]) => (
                <Fragment key={habit}>
                    <h2 className="whitespace-nowrap self-center">{habit}</h2>

                    {Array.from({ length: daysInMonth }).map((_, dayIndex) => {
                        const jot = jots.find(
                            (j) => new Date(j.date).getDate() === dayIndex + 1,
                        );
                        return jot ? (
                            <Checkbox
                                key={jot.id}
                                className={`aspect-square w-full min-w-4 h-full border-red-700 ${
                                    jot.isCompleted
                                        ? "bg-green-500 border-green-700"
                                        : "bg-gray-200 border-gray-400"
                                }`}
                                title={new Date(jot.date).toLocaleDateString()}
                                onClick={(e) => handleUpdateJot(e, jot)}
                            />
                        ) : (
                            <div key={`empty-${habit}-${dayIndex}`} />
                        );
                    })}
                </Fragment>
            ))}
        </section>
    );

    // return (
    //     <section className="flex flex-row gap-2 p-4 ">
    //         {Object.entries(jots).map(([habit, jots]) => (
    //             <div className="flex flex-col" key={habit}>
    //                 <input
    //                     className="text-sm font-semibold mb-1 pb-2 w-8 overflow-visible -rotate-45"
    //                     value={habit}
    //                 ></input>
    //                 <div className="flex flex-col">
    //                     {jots.map((jot) => (
    //                         <Checkbox
    //                             key={jot.id}
    //                             className={`w-5 h-5 rounded-sm border ${
    //                                 jot.isCompleted
    //                                     ? "bg-green-500 border-green-700"
    //                                     : "bg-gray-200 border-gray-400"
    //                             }`}
    //                             title={new Date(jot.date).toDateString()}
    //                             onClick={(e) => handleUpdateJot(e, jot)}
    //                         />
    //                     ))}
    //                 </div>
    //             </div>
    //         ))}
    //     </section>
    // );
}
