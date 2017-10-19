<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateProjectsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('projects', function (Blueprint $table) {
            $table->smallIncrements('id');
            $table->string('image_path', 255);
            $table->string('name', 30);
            $table->string('url', 255);
            $table->string('description', 500);
            $table->unsignedSmallInteger('heading');

            $table->foreign('heading')
                ->references('id')
                ->on('project_headings');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('project');
    }
}
