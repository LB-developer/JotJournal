"use client";
import { useEstablishUser } from "@/hooks/user";
import { updateJotCompletion } from "@/lib/jots/updateJotCompletion";
import { Jot, JotActionBody, JotCollection } from "@/types/jotTypes";
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
    const todayAsDate = new Date();

    const handleJotAction = async (
        e: React.MouseEvent<HTMLButtonElement | HTMLHeadingElement, MouseEvent>,
        jot: Jot,
        action: "update" | "delete",
    ): Promise<void> => {
        e.preventDefault();

        const jotID = jot.id;
        const habit = jot.habit;

        // TODO: create/delete tag for caching
        //   structure: ["id-habit1", "id-habit2", etc...] for as many habits that are in jots
        //
        // const tags: string[] = []
        // for (const key of Object.keys(jots)) {
        //   tags.push(`tag-${context.user.ID}-${key}`);
        // }

        let method: "PATCH" | "DELETE";
        let body: JotActionBody | null;
        let query = "";
        const headers = new Headers();
        headers.set("Content-Type", "application/json");

        switch (action) {
            case "update":
                method = "PATCH";
                const changedCompletion = !jot.isCompleted;
                body = { jotID, isCompleted: changedCompletion };
                break;
            case "delete":
                method = "DELETE";
                body = null;
                query = `?habit=${habit}&month=${month}&year=${year}`;
                break;
            default:
                throw new Error(`Invalid action: ${action}`);
        }

        // update/delete jot in db
        await fetch(`/api/jots` + query, {
            method,
            headers: headers,
            body: JSON.stringify(body),
        });

        // update/delete jot locally
        const updatedJots = { ...jots };

        if (action === "delete") {
            delete updatedJots[habit];
        } else {
            const length = updatedJots[habit].length - 1;
            const startingLeftNumber = 0;
            updateJotCompletion(
                updatedJots[habit],
                startingLeftNumber,
                length,
                jotID,
            );
        }

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
                        Number(year),
                        Number(month) - 1,
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
                    <h2
                        className="whitespace-nowrap self-center"
                        onClick={(e) => handleJotAction(e, jots[0], "delete")}
                    >
                        {habit}
                    </h2>

                    {Array.from({ length: daysInMonth }).map((_, dayIndex) => {
                        const jot = jots.find(
                            (j) => new Date(j.date).getDate() === dayIndex + 1,
                        );

                        const futureDate =
                            jot && new Date(jot?.date) > todayAsDate;

                        return jot ? (
                            <Checkbox
                                key={jot.id}
                                className={`aspect-square w-full min-w-4 h-full border-red-700 ${
                                    jot.isCompleted
                                        ? "bg-green-500 border-green-700"
                                        : "bg-gray-200 border-gray-400"
                                } 
                                ${futureDate && "opacity-50"}`}
                                title={new Date(jot.date).toLocaleDateString()}
                                onClick={(e) =>
                                    handleJotAction(e, jot, "update")
                                }
                                disabled={
                                    // disable if date is in the future
                                    futureDate
                                }
                            />
                        ) : (
                            <div key={`empty-${habit}-${dayIndex}`} />
                        );
                    })}
                </Fragment>
            ))}
        </section>
    );
}
